export interface Chat {
  id: string;
  generalUserId: string;
  professionalUserId: string;
  createdAt: string;
}

export interface Message {
  id: string;
  chatId: string;
  senderId: string;
  sentAt: string;
  type: string; // "StandardMessage", "QuestionMessage", "AnswerMessage", "PromotionalMessage"
  content: string | null;
  tags?: string[];
  questionId?: string;
  title?: string;
  body?: string;
  actionUrl?: string;
  imageUrl?: string;
}