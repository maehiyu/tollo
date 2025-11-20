import type { Message } from '../types';

interface MessageListProps {
  messages: Message[];
  currentUserId: string | null;
}

export const MessageList = ({ messages, currentUserId }: MessageListProps) => {
  return (
    <div
      className="message-list"
      style={{
        display: 'flex',
        flexDirection: 'column',
        gap: '12px',
        padding: '16px',
        overflowY: 'auto',
        flex: 1,
      }}
    >
      {messages.length === 0 ? (
        <div style={{ textAlign: 'center', color: '#999', padding: '32px' }}>
          No messages yet. Start the conversation!
        </div>
      ) : (
        messages.map((message) => {
          const isOwnMessage = message.senderId === currentUserId;
          return (
            <div
              key={message.id}
              style={{
                display: 'flex',
                justifyContent: isOwnMessage ? 'flex-end' : 'flex-start',
              }}
            >
              <div
                style={{
                  maxWidth: '70%',
                  padding: '12px 16px',
                  borderRadius: '12px',
                  backgroundColor: isOwnMessage ? '#0084ff' : '#e4e6eb',
                  color: isOwnMessage ? '#fff' : '#000',
                }}
              >
                <div style={{ fontSize: '14px', wordBreak: 'break-word' }}>
                  {message.content || message.title || 'Unknown message type'}
                </div>
                {message.type !== 'StandardMessage' && (
                  <div style={{ fontSize: '11px', marginTop: '4px', opacity: 0.8 }}>
                    {message.type}
                  </div>
                )}
                <div style={{ fontSize: '10px', marginTop: '4px', opacity: 0.7 }}>
                  {new Date(message.sentAt).toLocaleTimeString()}
                </div>
              </div>
            </div>
          );
        })
      )}
    </div>
  );
};