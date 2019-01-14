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

  // List<Post> getPostList() => List.from(this._postList);
  getPostList() async {
    var response = await _http.get(url);
    return json.decode(response.body);
  }

  Stream<List<Post>> get getPostUpdateListener => _postUpdated.stream;

  addPost(Post post) {
    _postList.add(post);
    _postUpdated.add(_postList.toList());
  }
}
