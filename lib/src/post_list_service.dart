import 'dart:async';
import 'dart:convert';
import 'package:angular/core.dart';
import 'package:http/browser_client.dart';

import './post_model.dart';

@Injectable()
class PostListService {
  final BrowserClient _http;
  final baseurl = 'http://localhost:3000/api/posts';
  static final _headers = {'Content-Type': 'application/json; charset=UTF-8'};
  Post post;
  List<Post> _postList = [];
  final StreamController<List<Post>> _postUpdated =
      new StreamController<List<Post>>.broadcast();

  PostListService(this._http);

  getPost(String postId) async {
    final url = '$baseurl/$postId';
    final response = await _http.get(url);
    return Post.fromJson(json.decode(response.body));
  }

  getAllPosts() async {
    final response = await _http.get(baseurl);
    final results = json.decode(response.body) as List;
    if (results != null) {
      _postList = results.map((json) => Post.fromJson(json)).toList();
      _postUpdated.add(_postList.toList());
    } else {
      _postUpdated.add(_postList.toList());
    }
  }

  Stream<List<Post>> get getPostUpdateListener => _postUpdated.stream;

  createPost(Post post) async {
    final response =
        await _http.post(baseurl, headers: _headers, body: json.encode(post));
    post = Post.fromJson(json.decode(response.body));
    _postList.add(post);
    _postUpdated.add(_postList.toList());
  }

  updatePost(Post post) async {
    final updateUrl = '$baseurl/${post.id}';
    await _http.put(updateUrl, headers: _headers, body: json.encode(post));
  }

  deletePost(String postId) async {
    final deleteUrl = '$baseurl/$postId';
    await _http.delete(deleteUrl, headers: _headers);
    _postList.removeWhere((post) => post.id == postId);
    _postUpdated.add(_postList.toList());
  }
}
