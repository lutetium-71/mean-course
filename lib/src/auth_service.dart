import 'dart:convert';
import 'package:angular/core.dart';
import 'package:http/browser_client.dart';

import './user_model.dart';

@Injectable()
class AuthService {
  final BrowserClient _http;
  String baseurl = 'http://localhost:3000/api/user';
  static final _headers = {'Content-Type': 'application/json; charset=UTF-8'};

  AuthService(this._http);

  createUser(User user) async {
    final url = '$baseurl/signup';
    final response =
        await _http.post(url, headers: _headers, body: json.encode(user));
    print(response.body);
  }

  loginUser(User user) async {
    final url = '$baseurl/login';
    final response =
        await _http.post(url, headers: _headers, body: json.encode(user));
    print(response.body);
  }

  // getPost(String postId) async {
  //   final getUrl = '$url/$postId';
  //   final response = await _http.get(getUrl);
  //   return Post.fromJson(json.decode(response.body));
  // }

  // getAllPosts() async {
  //   final response = await _http.get(url);
  //   final results = json.decode(response.body) as List;
  //   if (results != null) {
  //     _postList = results.map((json) => Post.fromJson(json)).toList();
  //     _postUpdated.add(_postList.toList());
  //   } else {
  //     _postUpdated.add(_postList.toList());
  //   }
  // }

  // Stream<List<Post>> get getPostUpdateListener => _postUpdated.stream;

  // updatePost(Post post) async {
  //   final updateUrl = '$url/${post.id}';
  //   await _http.put(updateUrl, headers: _headers, body: json.encode(post));
  // }

  // deletePost(String postId) async {
  //   final deleteUrl = '$url/$postId';
  //   await _http.delete(deleteUrl, headers: _headers);
  //   _postList.removeWhere((post) => post.id == postId);
  //   _postUpdated.add(_postList.toList());
  // }
}
