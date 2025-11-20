import { useState } from 'react';

interface MessageInputProps {
  onSend: (content: string) => void;
  disabled?: boolean;
}

export const MessageInput = ({ onSend, disabled }: MessageInputProps) => {
  const [content, setContent] = useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (content.trim() && !disabled) {
      onSend(content);
      setContent('');
    }
  };

  return (
    <form
      onSubmit={handleSubmit}
      style={{
        display: 'flex',
        gap: '8px',
        padding: '16px',
        borderTop: '1px solid #e0e0e0',
        backgroundColor: '#fff',
      }}
    >
      <input
        type="text"
        value={content}
        onChange={(e) => setContent(e.target.value)}
        placeholder="Type a message..."
        disabled={disabled}
        style={{
          flex: 1,
          padding: '12px 16px',
          border: '1px solid #e0e0e0',
          borderRadius: '24px',
          fontSize: '14px',
          outline: 'none',
        }}
      />
      <button
        type="submit"
        disabled={!content.trim() || disabled}
        style={{
          padding: '12px 24px',
          backgroundColor: content.trim() && !disabled ? '#0084ff' : '#e0e0e0',
          color: '#fff',
          border: 'none',
          borderRadius: '24px',
          fontSize: '14px',
          fontWeight: 600,
          cursor: content.trim() && !disabled ? 'pointer' : 'not-allowed',
          transition: 'background-color 0.2s',
        }}
      >
        Send
      </button>
    </form>
  );
};