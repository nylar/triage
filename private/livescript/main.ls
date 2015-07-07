Ticket = {list: ->
	m.request {
		method: 'GET'
		url: '/api/tickets/'
	}
}

App = {
	controller: ->
		tickets = Ticket.list!
		{
			tickets: tickets
		}
	view: (ctrl) ->
		m 'div', [(ctrl.tickets!.map ((ticket) -> m 'p', ticket.subject))]
}

m.mount (document.getElementById 'tickets'), App

