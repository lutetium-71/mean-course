import 'dart:async';
import 'dart:convert';
import 'package:angular/core.dart';
import 'package:http/browser_client.dart';

import './post_model.dart';

@Injectable()
class PostListService {
  final BrowserClient _http;
  String url = 'http://localhost:3000/api/posts';
  List<Post> _postList = [];
  final StreamController<List<Post>> _postUpdated =
      new StreamController<List<Post>>();

  PostListService(this._http);

  getPostList() async {
    final response = await _http.get(url);
    final _postList = (json.decode(response.body) as List)
        .map((json) => Post.fromJson(json))
        .toList();
    _postUpdated.add(_postList.toList());
  }

  Stream<List<Post>> get getPostUpdateListener => _postUpdated.stream;

  addPost(Post post) {
    _postList.add(post);
    _postUpdated.add(_postList.toList());
  }
}
