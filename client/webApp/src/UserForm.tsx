import { useState } from 'react';
import { createUser, getUserByEmail, JsUser } from 'shared';
import './UserForm.css';

// Define the type for the created user based on the GraphQL schema
// This helps with type safety and autocompletion.
type CreatedUser = JsUser;

function UserForm() {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [description, setDescription] = useState('');

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [createdUser, setCreatedUser] = useState<CreatedUser | null>(null);

  const [profileType, setProfileType] = useState<'none' | 'general' | 'professional'>('none');

  // General Profile states
  const [generalPoints, setGeneralPoints] = useState('');
  const [generalIntroduction, setGeneralIntroduction] = useState('');

  // Professional Profile states
  const [proBadgeUrl, setProBadgeUrl] = useState('');
  const [proBiography, setProBiography] = useState('');

  // Search states
  const [searchEmail, setSearchEmail] = useState('');
  const [searchedUser, setSearchedUser] = useState<CreatedUser | null>(null);
  const [searchLoading, setSearchLoading] = useState(false);
  const [searchError, setSearchError] = useState<string | null>(null);

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    setLoading(true);
    setError(null);
    setCreatedUser(null);

    try {
      let generalProfile: { points: number; introduction: string } | null = null;
      let professionalProfile: { proBadgeUrl: string; biography: string } | null = null;

      if (profileType === 'general') {
        generalProfile = {
          points: parseInt(generalPoints, 10),
          introduction: generalIntroduction,
        };
      } else if (profileType === 'professional') {
        professionalProfile = {
          proBadgeUrl: proBadgeUrl,
          biography: proBiography,
        };
      }

      const user = await createUser(
        name,
        email,
        description || null,
        generalProfile,
        professionalProfile
      );
      console.log('Received user from createUser:', user);
      if (user) {
        setCreatedUser(user);
        // Clear form on success
        setName('');
        setEmail('');
        setDescription('');
        setProfileType('none');
        setGeneralPoints('');
        setGeneralIntroduction('');
        setProBadgeUrl('');
        setProBiography('');
      } else {
        setError('Failed to create user. The server returned no user data.');
      }
    } catch (e: any) {
      console.error('[React] Error creating user:', e);
      setError(e.message || 'An unknown error occurred.');
    } finally {
      setLoading(false);
    }
  };

  const handleSearch = async (event: React.FormEvent) => {
    event.preventDefault();
    setSearchLoading(true);
    setSearchError(null);
    setSearchedUser(null);

    try {
      const user = await getUserByEmail(searchEmail);
      console.log('Received user from getUserByEmail:', user);
      if (user) {
        setSearchedUser(user);
      } else {
        setSearchError('User not found with this email.');
      }
    } catch (e: any) {
      console.error('[React] Error searching user:', e);
      setSearchError(e.message || 'An unknown error occurred during search.');
    } finally {
      setSearchLoading(false);
    }
  };

  return (
    <div className="user-form-container">
      <h2>Create a New User</h2>
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="name">Name:</label>
          <input
            id="name"
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
        </div>
        <div className="form-group">
          <label htmlFor="email">Email:</label>
          <input
            id="email"
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>
        <div className="form-group">
          <label htmlFor="description">Description (Optional):</label>
          <textarea
            id="description"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
          />
        </div>

        <div className="form-group">
          <label>Profile Type:</label>
          <div>
            <label>
              <input
                type="radio"
                value="none"
                checked={profileType === 'none'}
                onChange={() => setProfileType('none')}
              />{' '}
              None
            </label>
            <label>
              <input
                type="radio"
                value="general"
                checked={profileType === 'general'}
                onChange={() => setProfileType('general')}
              />{' '}
              General Profile
            </label>
            <label>
              <input
                type="radio"
                value="professional"
                checked={profileType === 'professional'}
                onChange={() => setProfileType('professional')}
              />{' '}
              Professional Profile
            </label>
          </div>
        </div>

        {profileType === 'general' && (
          <div className="profile-section">
            <h3>General Profile</h3>
            <div className="form-group">
              <label htmlFor="generalPoints">Points:</label>
              <input
                id="generalPoints"
                type="number"
                value={generalPoints}
                onChange={(e) => setGeneralPoints(e.target.value)}
                required
              />
            </div>
            <div className="form-group">
              <label htmlFor="generalIntroduction">Introduction:</label>
              <textarea
                id="generalIntroduction"
                value={generalIntroduction}
                onChange={(e) => setGeneralIntroduction(e.target.value)}
                required
              />
            </div>
          </div>
        )}

        {profileType === 'professional' && (
          <div className="profile-section">
            <h3>Professional Profile</h3>
            <div className="form-group">
              <label htmlFor="proBadgeUrl">Pro Badge URL:</label>
              <input
                id="proBadgeUrl"
                type="text"
                value={proBadgeUrl}
                onChange={(e) => setProBadgeUrl(e.target.value)}
                required
              />
            </div>
            <div className="form-group">
              <label htmlFor="proBiography">Biography:</label>
              <textarea
                id="proBiography"
                value={proBiography}
                onChange={(e) => setProBiography(e.target.value)}
                required
              />
            </div>
          </div>
        )}

        <button type="submit" disabled={loading}>
          {loading ? 'Creating...' : 'Create User'}
        </button>
      </form>

      {error && <p className="error-message">Error: {error}</p>}
      
      {createdUser && (
        <div className="success-message">
          <h3>User Created Successfully!</h3>
          <p><strong>ID:</strong> {createdUser.id}</p>
          <p><strong>Name:</strong> {createdUser.name}</p>
          <p><strong>Email:</strong> {createdUser.email}</p>
        </div>
      )}

      <hr />

      <h2>Search User by Email</h2>
      <form onSubmit={handleSearch}>
        <div className="form-group">
          <label htmlFor="searchEmail">Email:</label>
          <input
            id="searchEmail"
            type="email"
            value={searchEmail}
            onChange={(e) => setSearchEmail(e.target.value)}
            required
          />
        </div>
        <button type="submit" disabled={searchLoading}>
          {searchLoading ? 'Searching...' : 'Search User'}
        </button>
      </form>

      {searchError && <p className="error-message">Error: {searchError}</p>}

      {searchedUser && (
        <div className="success-message">
          <h3>User Found!</h3>
          <p><strong>ID:</strong> {searchedUser.id}</p>
          <p><strong>Name:</strong> {searchedUser.name}</p>
          <p><strong>Email:</strong> {searchedUser.email}</p>
        </div>
      )}
    </div>
  );
}

export default UserForm;
