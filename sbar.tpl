<div>
	<form action="" method="POST" enctype="multipart/form-data">
	<input type="search" name="searching_query" placeholder="Search..." required><br>
	<input type="submit" name="searching_preview" value="Search" style="display:none">
	<input type="checkbox" id="search_this_directory" name="search_app_search_directory" value="search_app_search_directory">
	<label for="search_this_directory">Include filenames?</label>
	</form>
</div>

% if(! ~ $#post_arg_search_app_search_directory 0){
%	filenames=`{walk sites/$SERVER_NAME | grep -v -e '_werc' -e '.meta' | grep -i $post_arg_searching_query}
	<p>Filenames:
%		for(match in $filenames){
%			href_path=`{echo $match | sed -e 's/.md//g' -e 's|sites/'$SERVER_NAME'||g'}
			<a href="%($href_path%)">%($href_path%)</a>
%		}
%		if(~ $#filenames 0 || ){
			<p>No filename matches for "<i>%($"post_arg_searching_query%)</i>"</p>
%		}
% }

% if (! ~ $"post_arg_searching_preview) {
%	tank $post_arg_searching_query 
% }


