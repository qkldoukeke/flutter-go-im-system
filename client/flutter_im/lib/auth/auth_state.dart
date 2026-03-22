// auth_state.dart

abstract class AuthState {}

class AuthInitial extends AuthState {}

class AuthAuthenticated extends AuthState {
  final String username;
  AuthAuthenticated(this.username);
}

class AuthUnauthenticated extends AuthState {}