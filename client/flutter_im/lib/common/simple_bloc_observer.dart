// simple_bloc_observer.dart

import 'package:flutter_bloc/flutter_bloc.dart';

class SimpleBlocObserver extends BlocObserver {
  @override
  void onEvent(Bloc bloc, Object? event) {
    super.onEvent(bloc, event);
    print('[$bloc] event: $event');
  }

  @override
  void onChange(Bloc bloc, Change change) {
    super.onChange(bloc, change);
    print('[$bloc] change: $change');
  }

  @override
  void onError(Bloc bloc, Object? error, StackTrace stackTrace) {
    super.onError(bloc, error, stackTrace);
    print('[$bloc] error: $error');
  }

  @override
  void onTransition(Bloc bloc, Transition transition) {
    super.onTransition(bloc, transition);
    print('[$bloc] transition: $transition');
  }
}