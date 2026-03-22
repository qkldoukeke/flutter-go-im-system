// chat_page.dart

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'chat_bloc.dart';
import 'chat_event.dart';
import 'chat_state.dart';

class ChatPage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Chat'),
      ),
      body: BlocListener<ChatBloc, ChatState>(
        listener: (context, state) {
          if (state is ChatMessageReceivedState) {
            // Handle received messages
          }
        },
        child: Column(
          children: <Widget>[
            Expanded(child: Container()),
            TextField(onSubmitted: (message) {
              BlocProvider.of<ChatBloc>(context).add(ChatMessageSent(message));
            }),
          ],
        ),
      ),
    );
  }
}