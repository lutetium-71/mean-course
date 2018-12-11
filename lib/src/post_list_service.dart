import 'dart:async';
import 'package:angular/core.dart';

import './post_model.dart';

@Injectable()
class PostListService {
  List<Post> _postList = <Post>[];
  final StreamController<List<Post>> _postUpdated =
      new StreamController<List<Post>>();

  List<Post> getPostList() => List.from(this._postList);

  StreamController<List<Post>> getPostUpdateListener() => this._postUpdated;

  addPost(Post post) {
    this._postList.add(post);
    this._postUpdated.add(List.from(this._postList));
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
