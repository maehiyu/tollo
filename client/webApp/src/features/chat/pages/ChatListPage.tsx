import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { getCurrentUserId } from 'shared';
import { useChat } from '../hooks/useChat';
import { ChatListItem } from '../components/ChatListItem';
import type { Chat } from '../types';

export const ChatListPage = () => {
  const navigate = useNavigate();
  const { fetchUserChats, createNewChat, loading, error } = useChat();
  const [chats, setChats] = useState<Chat[]>([]);
  const [currentUserId, setCurrentUserId] = useState<string | null>(null);
  const [showCreateForm, setShowCreateForm] = useState(false);
  const [professionalUserId, setProfessionalUserId] = useState('');

  useEffect(() => {
    const userId = getCurrentUserId();
    setCurrentUserId(userId);

    if (userId) {
      fetchUserChats(userId).then((fetchedChats) => {
        if (fetchedChats) {
          setChats(fetchedChats);
        }
      });
    }
  }, []);

  const handleChatClick = (chatId: string) => {
    navigate(`/chat/${chatId}`);
  };

  const handleCreateChat = async () => {
    if (!currentUserId || !professionalUserId.trim()) return;

    const newChat = await createNewChat(currentUserId, professionalUserId.trim());
    if (newChat) {
      setChats((prev) => [newChat, ...prev]);
      setShowCreateForm(false);
      setProfessionalUserId('');
      navigate(`/chat/${newChat.id}`);
    }
  };

  return (
    <div style={{ maxWidth: '800px', margin: '0 auto', padding: '24px' }}>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '24px' }}>
        <h1 style={{ margin: 0 }}>Chats</h1>
        <button
          onClick={() => setShowCreateForm(!showCreateForm)}
          style={{
            padding: '10px 20px',
            backgroundColor: '#0084ff',
            color: '#fff',
            border: 'none',
            borderRadius: '8px',
            fontSize: '14px',
            fontWeight: 600,
            cursor: 'pointer',
          }}
        >
          {showCreateForm ? 'Cancel' : '+ New Chat'}
        </button>
      </div>

      {showCreateForm && (
        <div
          style={{
            backgroundColor: '#fff',
            border: '1px solid #e0e0e0',
            borderRadius: '8px',
            padding: '20px',
            marginBottom: '16px',
          }}
        >
          <h3 style={{ margin: '0 0 16px 0' }}>Create New Chat</h3>
          <div style={{ marginBottom: '12px' }}>
            <label style={{ display: 'block', marginBottom: '8px', fontSize: '14px', fontWeight: 500 }}>
              Professional User ID:
            </label>
            <input
              type="text"
              value={professionalUserId}
              onChange={(e) => setProfessionalUserId(e.target.value)}
              placeholder="Enter professional user ID"
              style={{
                width: '100%',
                padding: '10px 12px',
                border: '1px solid #e0e0e0',
                borderRadius: '6px',
                fontSize: '14px',
              }}
            />
          </div>
          <button
            onClick={handleCreateChat}
            disabled={!professionalUserId.trim() || loading}
            style={{
              padding: '10px 20px',
              backgroundColor: professionalUserId.trim() && !loading ? '#0084ff' : '#e0e0e0',
              color: '#fff',
              border: 'none',
              borderRadius: '6px',
              fontSize: '14px',
              fontWeight: 600,
              cursor: professionalUserId.trim() && !loading ? 'pointer' : 'not-allowed',
            }}
          >
            {loading ? 'Creating...' : 'Create Chat'}
          </button>
        </div>
      )}

      {!currentUserId && (
        <div style={{ padding: '16px', backgroundColor: '#fff3cd', borderRadius: '8px', marginBottom: '16px' }}>
          Please log in to view your chats.
        </div>
      )}

      {error && (
        <div style={{ padding: '16px', backgroundColor: '#f8d7da', borderRadius: '8px', marginBottom: '16px' }}>
          Error: {error}
        </div>
      )}

      {loading ? (
        <div style={{ textAlign: 'center', padding: '32px', color: '#999' }}>
          Loading chats...
        </div>
      ) : chats.length === 0 ? (
        <div style={{ textAlign: 'center', padding: '32px', color: '#999' }}>
          No chats yet. Start a new conversation!
        </div>
      ) : (
        <div style={{ backgroundColor: '#fff', borderRadius: '8px', border: '1px solid #e0e0e0' }}>
          {chats.map((chat) => (
            <ChatListItem key={chat.id} chat={chat} onClick={handleChatClick} />
          ))}
        </div>
      )}
    </div>
  );
};