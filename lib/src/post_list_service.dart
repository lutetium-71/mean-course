import 'dart:async';
import 'package:angular/core.dart';

import './post_model.dart';

@Injectable()
class PostListService {
  List<Post> _postList = <Post>[];
  final StreamController<List<Post>> _postUpdated =
      new StreamController<List<Post>>();

  List<Post> getPostList() => List.from(this._postList);

  Stream<List<Post>> get getPostUpdateListener => _postUpdated.stream;

  addPost(Post post) {
    this._postList.add(post);
    this._postUpdated.add(List.from(this._postList));
  }
}
