import type { Chat } from '../types';

interface ChatListItemProps {
  chat: Chat;
  onClick: (chatId: string) => void;
}

export const ChatListItem = ({ chat, onClick }: ChatListItemProps) => {
  return (
    <div
      className="chat-list-item"
      onClick={() => onClick(chat.id)}
      style={{
        padding: '16px',
        borderBottom: '1px solid #e0e0e0',
        cursor: 'pointer',
        transition: 'background-color 0.2s',
      }}
      onMouseEnter={(e) => {
        e.currentTarget.style.backgroundColor = '#f5f5f5';
      }}
      onMouseLeave={(e) => {
        e.currentTarget.style.backgroundColor = 'transparent';
      }}
    >
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
        <div>
          <h3 style={{ margin: 0, fontSize: '16px', fontWeight: 600 }}>Chat {chat.id.slice(0, 8)}</h3>
          <p style={{ margin: '4px 0 0', fontSize: '14px', color: '#666' }}>
            Users: {chat.generalUserId.slice(0, 8)} & {chat.professionalUserId.slice(0, 8)}
          </p>
        </div>
        <div style={{ fontSize: '12px', color: '#999' }}>
          {new Date(chat.createdAt).toLocaleDateString()}
        </div>
      </div>
    </div>
  );
};