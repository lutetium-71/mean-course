import 'dart:async';
import 'package:angular/core.dart';

import './post_model.dart';

@Injectable()
class PostListService {
  List<Post> postList = <Post>[];
  final _postUpdated = new StreamController<List<Post>>();

  Future<List<Post>> getPostList() async => List.from(this.postList);

  Stream<List<Post>> getPostUpdateListener() => this._postUpdated.stream;

  addPost(Post post) {
    this.postList.add(post);
    this._postUpdated.add(List.from(this.postList));
  }

  //   PostCreateComponent(this.postListService);

  // void onAddPost(NgForm form) {
  //   Post post = Post(form.value["title"], form.value["content"]);
  //   // _postCreated.add(post);
  //   this.postListService.addPost(post);
  // }

  // bool checkEntry() => enteredTitle.isEmpty || enteredContent.isEmpty;

  // // @Output()
  // // Stream<Post> get postCreated => _postCreated.stream;
}
