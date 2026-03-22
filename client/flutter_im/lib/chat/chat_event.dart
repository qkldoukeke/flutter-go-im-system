// chat_event.dart

abstract class ChatEvent {}

class ChatMessageSent extends ChatEvent {
  final String message;
  ChatMessageSent(this.message);
}

class ChatMessageReceived extends ChatEvent {
  final String message;
  ChatMessageReceived(this.message);
}