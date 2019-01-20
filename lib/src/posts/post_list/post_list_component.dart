import 'dart:async';

import 'package:angular/angular.dart';
import 'package:angular_router/angular_router.dart';

import 'package:angular_components/material_expansionpanel/material_expansionpanel.dart';
import '../../route_paths.dart';
import '../../post_model.dart';
import '../../post_list_service.dart';

@Component(
  selector: 'app-post-list',
  styleUrls: ['post_list_component.css'],
  templateUrl: 'post_list_component.html',
  directives: [
    MaterialExpansionPanel,
    NgFor,
    NgIf,
    coreDirectives,
  ],
)
class PostListComponent implements OnDestroy, OnActivate {
  final PostListService _postListService;
  final Router _router;
  StreamSubscription _postsSubscription;
  List<Post> postList = [];

  PostListComponent(this._postListService, this._router);

  @override
  void onActivate(_, RouterState current) {
    _postListService.getAllPosts();
    _postsSubscription = _postListService.getPostUpdateListener
        .listen((List<Post> posts) => postList = posts);
  }

  deletePost(String id) {
    _postListService.deletePost(id);
  }

  getPost(String id) {
    _gotoPost(id);
  }

  Future<NavigationResult> _gotoPost(String id) =>
      _router.navigate(_postUrl(id));

  // Future<NavigationResult> goBack() => _router.navigate(
  //     RoutePaths.posts.toUrl(),
  //     NavigationParams(queryParameters: {id: '${post.id}'}));

  String _postUrl(String id) =>
      RoutePaths.edit.toUrl(parameters: {postId: '$id'});

  ngOnDestroy() {
    _postsSubscription.cancel();
  }
}
