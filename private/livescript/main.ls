'use strict'

(->
  app = app || {}
  app.TicketList = ->
    m.request {
      method: 'GET'
      url: '/api/tickets'
    }
  app.vm = {}
  app.vm.init = -> @tickets = new app.TicketList
  app.controller = -> app.vm.init!
  app.view = -> app.vm.tickets!.map ((ticket) -> m 'p', ticket.subject)
  m.mount (document.getElementById 'tickets'), app)!