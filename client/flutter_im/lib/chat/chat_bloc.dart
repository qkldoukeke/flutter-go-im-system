// chat_bloc.dart

import 'package:flutter_bloc/flutter_bloc.dart';
import 'chat_event.dart';
import 'chat_state.dart';

class ChatBloc extends Bloc<ChatEvent, ChatState> {
  ChatBloc() : super(ChatInitial());

  @override
  Stream<ChatState> mapEventToState(ChatEvent event) async* {
    if (event is ChatMessageSent) {
      yield ChatMessageSentState(event.message);
    } else if (event is ChatMessageReceived) {
      yield ChatMessageReceivedState(event.message);
    }
  }
}