import React from 'react'
import ReactDOM from 'react-dom'
import App from './App'
import store from './redux'
import { Provider } from 'react-redux'
import reportWebVitals from './reportWebVitals'
import { Provider as BusProvider } from '@/hooks/useBus'
import 'antd/dist/antd.css'
import '@/assets/style/normalize.css'
import '@/assets/style/reset.less'

ReactDOM.render(
  <BusProvider>
    <Provider store={store}>
      <App/>
    </Provider>
  </BusProvider>,
  document.getElementById('root')
)

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals()
