import thunk from 'redux-thunk'
import { compose, applyMiddleware, createStore } from 'redux'
import { composeWithDevTools } from 'redux-devtools-extension'
import rootReducer from './rootReducers'

let storeEnhancers

if (process.env.NODE_ENV === 'production') {
  storeEnhancers = compose(applyMiddleware(thunk))
} else {
  storeEnhancers = compose(composeWithDevTools(applyMiddleware(thunk)))
}

const configureStore = (initialState = {}) => {
  const store = createStore(rootReducer, initialState, storeEnhancers)
  console.log('module', module)
  // Enable Webpack hot module replacement for reducer
  if (module.hot && process.env.NODE_ENV === 'production') {
    module.hot.accept('./rootReduces', () => {
      console.log('replacing reducer...')
      const nextRootReducer = require('./rootReducers').default
      store.replaceReducer(nextRootReducer)
    })
  }
  return store
}


export default configureStore()
