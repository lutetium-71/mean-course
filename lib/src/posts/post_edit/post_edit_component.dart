import 'dart:async';

import 'package:angular/angular.dart';
import 'package:angular_forms/angular_forms.dart';
import 'package:angular_router/angular_router.dart';
import 'package:angular_components/focus/focus.dart';
import 'package:angular_components/material_button/material_button.dart';
import 'package:angular_components/material_icon/material_icon.dart';
import 'package:angular_components/material_input/material_input.dart';

import '../../post_model.dart';
import '../../post_list_service.dart';

@Component(
  selector: 'app-post-edit',
  styleUrls: [
    'package:angular_components/css/mdc_web/card/mdc-card.scss.css',
    'post_edit_component.css'
  ],
  templateUrl: 'post_edit_component.html',
  directives: [
    AutoFocusDirective,
    formDirectives,
    MaterialButtonComponent,
    MaterialIconComponent,
    materialInputDirectives,
    NgForm,
    NgIf,
  ],
)
class PostEditComponent implements OnActivate, OnDestroy {
  final PostListService _postListService;
  final Router _router;
  StreamSubscription _postsSubscription;

  PostEditComponent(this._postListService, this._router);

  @override
  void onActivate(_, RouterState routerState) async {
    final id = getId(routerState.queryParameters);
    await _getPost(id);
  }

  Future<void> _getPost(String id) async {
    await _postListService.getPost(id);
  }

  onEditPost(NgForm form) {
    Post post = Post(null, form.value["title"], form.value["content"]);
    _postListService.updatePost(post);
  }

  ngOnDestroy() {
    _postsSubscription.cancel();
  }
}
