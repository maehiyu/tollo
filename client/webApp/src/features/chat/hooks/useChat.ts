import { useState } from 'react';
import { getUserChats, getChatMessages, sendStandardMessage, createChat } from 'shared';
import type { Chat, Message } from '../types';

export const useChat = () => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchUserChats = async (userId: string): Promise<Chat[] | null> => {
    setLoading(true);
    setError(null);
    try {
      const chats = await getUserChats(userId);
      return chats ? (chats as Chat[]) : null;
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to fetch chats');
      return null;
    } finally {
      setLoading(false);
    }
  };

  const fetchChatMessages = async (chatId: string): Promise<Message[] | null> => {
    setLoading(true);
    setError(null);
    try {
      const messages = await getChatMessages(chatId);
      return messages ? (messages as Message[]) : null;
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to fetch messages');
      return null;
    } finally {
      setLoading(false);
    }
  };

  const sendMessage = async (chatId: string, content: string): Promise<Message | null> => {
    setLoading(true);
    setError(null);
    try {
      const message = await sendStandardMessage(chatId, content);
      return message ? (message as Message) : null;
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to send message');
      return null;
    } finally {
      setLoading(false);
    }
  };

  const createNewChat = async (generalUserId: string, professionalUserId: string): Promise<Chat | null> => {
    setLoading(true);
    setError(null);
    try {
      const chat = await createChat(generalUserId, professionalUserId);
      return chat ? (chat as Chat) : null;
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to create chat');
      return null;
    } finally {
      setLoading(false);
    }
  };

  return {
    loading,
    error,
    fetchUserChats,
    fetchChatMessages,
    sendMessage,
    createNewChat,
  };
};