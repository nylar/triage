(function(){"use strict";!function(){var t;return t=t||{},t.TicketList=function(){return m.request({method:"GET",url:"/api/tickets"})},t.vm={},t.vm.init=function(){return this.tickets=new t.TicketList},t.controller=function(){return t.vm.init()},t.view=function(){return t.vm.tickets().map(function(t){return m("p",t.subject)})},m.mount(document.getElementById("tickets"),t)}()}).call(this);
