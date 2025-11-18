import { useState } from 'react'
import { login, signUp } from 'shared'
import './LoginPage.css'

interface LoginPageProps {
  onLogin: () => void
}

export default function LoginPage({ onLogin }: LoginPageProps) {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')
  const [mode, setMode] = useState<'login' | 'signup'>('login')

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()

    if (!email || !password) {
      setError('Email and password are required')
      return
    }

    // AuthService を呼び出し
    if (mode === 'signup') {
      const userId = signUp(email, password)
      if (!userId) {
        setError('User already exists')
        return
      }
    } else {
      const userId = login(email, password)
      if (!userId) {
        setError('Invalid email or password')
        return
      }
    }

    onLogin()
  }

  return (
    <div className="login-container">
      <h2>{mode === 'login' ? 'Login' : 'Sign Up'}</h2>
      <form onSubmit={handleSubmit}>
        {error && <div className="error">{error}</div>}

        <div className="form-group">
          <label htmlFor="email">Email:</label>
          <input
            id="email"
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="user@example.com"
          />
        </div>

        <div className="form-group">
          <label htmlFor="password">Password:</label>
          <input
            id="password"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="password"
          />
        </div>

        <button type="submit">{mode === 'login' ? 'Login' : 'Sign Up'}</button>
      </form>

      <div style={{ marginTop: '1rem' }}>
        {mode === 'login' ? (
          <span>
            Don't have an account?{' '}
            <button onClick={() => setMode('signup')} style={{ background: 'none', border: 'none', color: 'blue', textDecoration: 'underline', cursor: 'pointer' }}>
              Sign Up
            </button>
          </span>
        ) : (
          <span>
            Already have an account?{' '}
            <button onClick={() => setMode('login')} style={{ background: 'none', border: 'none', color: 'blue', textDecoration: 'underline', cursor: 'pointer' }}>
              Login
            </button>
          </span>
        )}
      </div>
    </div>
  )
}