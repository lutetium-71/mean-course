import 'package:angular_router/angular_router.dart';

const postId = 'id';

class RoutePaths {
  static final posts = RoutePath(path: '');
  static final create = RoutePath(path: 'create');
  static final edit = RoutePath(path: 'edit/:$postId');
}

String getId(Map<String, String> parameters) {
  final id = parameters[postId];
  return id == null ? null : id;
}
