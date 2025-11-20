import { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { getCurrentUserId } from 'shared';
import { useChat } from '../hooks/useChat';
import { MessageList } from '../components/MessageList';
import { MessageInput } from '../components/MessageInput';
import type { Message } from '../types';

export const ChatDetailPage = () => {
  const { chatId } = useParams<{ chatId: string }>();
  const navigate = useNavigate();
  const { fetchChatMessages, sendMessage, loading, error } = useChat();
  const [messages, setMessages] = useState<Message[]>([]);
  const [currentUserId, setCurrentUserId] = useState<string | null>(null);
  const [sending, setSending] = useState(false);

  useEffect(() => {
    const userId = getCurrentUserId();
    setCurrentUserId(userId);

    if (chatId) {
      loadMessages();
    }
  }, [chatId]);

  const loadMessages = async () => {
    if (!chatId) return;

    const fetchedMessages = await fetchChatMessages(chatId);
    if (fetchedMessages) {
      setMessages(fetchedMessages);
    }
  };

  const handleSendMessage = async (content: string) => {
    if (!chatId || sending) return;

    setSending(true);
    try {
      const newMessage = await sendMessage(chatId, content);
      if (newMessage) {
        setMessages((prev) => [...prev, newMessage]);
      }
    } finally {
      setSending(false);
    }
  };

  return (
    <div style={{ maxWidth: '800px', margin: '0 auto', height: '100vh', display: 'flex', flexDirection: 'column' }}>
      <div
        style={{
          padding: '16px 24px',
          borderBottom: '1px solid #e0e0e0',
          backgroundColor: '#fff',
          display: 'flex',
          alignItems: 'center',
          gap: '16px',
        }}
      >
        <button
          onClick={() => navigate('/chat')}
          style={{
            padding: '8px 16px',
            backgroundColor: '#f0f0f0',
            border: 'none',
            borderRadius: '4px',
            cursor: 'pointer',
          }}
        >
          ← Back
        </button>
        <h2 style={{ margin: 0 }}>Chat {chatId?.slice(0, 8)}</h2>
      </div>

      {error && (
        <div style={{ padding: '16px', backgroundColor: '#f8d7da', margin: '16px' }}>
          Error: {error}
        </div>
      )}

      {loading && messages.length === 0 ? (
        <div style={{ flex: 1, display: 'flex', alignItems: 'center', justifyContent: 'center', color: '#999' }}>
          Loading messages...
        </div>
      ) : (
        <>
          <MessageList messages={messages} currentUserId={currentUserId} />
          <MessageInput onSend={handleSendMessage} disabled={sending || !currentUserId} />
        </>
      )}
    </div>
  );
};