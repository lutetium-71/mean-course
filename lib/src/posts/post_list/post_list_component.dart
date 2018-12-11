import 'dart:async';
import 'package:angular/angular.dart';

import 'package:angular_components/material_expansionpanel/material_expansionpanel.dart';
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
  ],
  providers: [ClassProvider(PostListService)],
)
class PostListComponent implements OnInit, OnDestroy {
  final PostListService postListService;
  StreamSubscription _postsSubscription;
  List<Post> postList = [];

  PostListComponent(this.postListService);

  @override
  ngOnInit() {
    this.postList = this.postListService.getPostList();
    this._postsSubscription = this
        .postListService
        .getPostUpdateListener()
        .stream
        .listen((List<Post> posts) => this.postList = posts);
  }

  ngOnDestroy() {
    this._postsSubscription.cancel();
  }
}
