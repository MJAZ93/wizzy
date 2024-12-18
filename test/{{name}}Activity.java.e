@@//New Services comes here
// Servico {{name}} {{feature}}
fun {%if(type==GET)%} fetch{{name}} {%endif%} {%if(type==POST)%} post{{name}} {%endif%}(
	{%for (params as param)%}
		{{param}}
	{%endfor%}
) {
	return {{objectOut}};
}
-@@

@@//New Declaration Comes here
//{{name}} {{feature}}
-@@