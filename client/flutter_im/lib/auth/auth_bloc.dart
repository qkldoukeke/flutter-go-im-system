// auth_bloc.dart

import 'package:flutter_bloc/flutter_bloc.dart';
import 'auth_event.dart';
import 'auth_state.dart';

class AuthBloc extends Bloc<AuthEvent, AuthState> {
  AuthBloc() : super(AuthInitial());

  @override
  Stream<AuthState> mapEventToState(AuthEvent event) async* {
    if (event is AuthLoginRequested) {
      // Handle login
      yield AuthAuthenticated(event.username);
    } else if (event is AuthLogoutRequested) {
      // Handle logout
      yield AuthUnauthenticated();
    }
  }
}