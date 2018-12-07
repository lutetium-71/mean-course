import 'package:angular/angular.dart';
import 'package:angular_components/material_expansionpanel/material_expansionpanel.dart';
import 'package:angular_components/material_expansionpanel/material_expansionpanel_set.dart';
import '../../models/post.dart';

@Component(
    selector: 'app-post-list',
    styleUrls: ['post_list_component.css'],
    templateUrl: 'post_list_component.html',
    providers: [],
    directives: [
      MaterialExpansionPanel,
      MaterialExpansionPanelSet,
      NgFor,
    ])
class PostListComponent {
  List<Post> posts = [
    Post('First Post', 'First Post to be displayed'),
    Post('Second Post', 'Second Post to be displayed'),
    Post('Third Post', 'Third Post to be displayed')
  ];
}
