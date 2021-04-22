import routes from '@/routes'
import { BrowserRouter, Route, Switch } from 'react-router-dom'

const App = props => {
  function renderRoutes (routes, contextPath) {
    const children = []

    const renderRoute = (item, routeContextPath) => {
      let newContextPath = item.path ? `${routeContextPath}/${item.path}` : routeContextPath
      newContextPath = newContextPath.replace(/\/+/g, '/')
      if (!item.component) {
        return
      }
      if (item.childRoutes) {
        const childRoutes = renderRoutes(item.childRoutes, newContextPath)
        children.push(
          <Route
            key={newContextPath}
            render={props => <item.component {...props}>{childRoutes}</item.component>}
            path={newContextPath}
          />
        )
      } else {
        children.push(<Route key={newContextPath} component={item.component} path={newContextPath} exact/>)
      }
    }

    routes.forEach(item => renderRoute(item, contextPath))

    return <Switch>{children}</Switch>
  }

  const Routers = renderRoutes(routes, '/')
  return (
    <BrowserRouter>
      {Routers}
    </BrowserRouter>
  )
}

export default App
