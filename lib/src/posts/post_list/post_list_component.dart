import 'package:angular/angular.dart';
import 'package:angular_components/material_expansionpanel/material_expansionpanel.dart';
import 'package:angular_components/material_expansionpanel/material_expansionpanel_set.dart';

@Component(
    selector: 'app-post-list',
    styleUrls: ['post_list_component.css'],
    templateUrl: 'post_list_component.html',
    providers: [],
    directives: [
      MaterialExpansionPanel,
      MaterialExpansionPanelSet,
    ])
class PostListComponent {}
