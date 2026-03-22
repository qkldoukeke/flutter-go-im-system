// chat_state.dart

abstract class ChatState {}

class ChatInitial extends ChatState {}

class ChatMessageSentState extends ChatState {
  final String message;
  ChatMessageSentState(this.message);
}

class ChatMessageReceivedState extends ChatState {
  final String message;
  ChatMessageReceivedState(this.message);
}