(function(){var t,e;t={list:function(){return m.request({method:"GET",url:"/api/tickets/"})}},e={controller:function(){var e;return e=t.list(),{tickets:e}},view:function(t){return m("div",[t.tickets().map(function(t){return m("p",t.subject)})])}},m.mount(document.getElementById("tickets"),e)}).call(this);
