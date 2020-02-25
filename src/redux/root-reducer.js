import { combineReducers, createStore, applyMiddleware, compose  } from 'redux'
import { persistReducer, persistStore } from 'redux-persist'
import { localizeReducer } from 'react-localize-redux';
import { combineEpics, createEpicMiddleware } from 'redux-observable';
import autoMergeLevel2 from 'redux-persist/es/stateReconciler/autoMergeLevel2';
import { connectRouter, routerMiddleware } from 'connected-react-router';
import storage from 'redux-persist/lib/storage'
import product from './product/productReducer'
import { createBrowserHistory } from 'history';
import thunkMiddleware from 'redux-thunk'

export const browserHistory = createBrowserHistory();

export const rootEpic = combineEpics();

const epicMiddleware = createEpicMiddleware();

export const compareApp = combineReducers({
  localize: localizeReducer,
  router: connectRouter(browserHistory),
  product
});

const persistConfig = {
  key: 'root',
  stateReconciler: autoMergeLevel2,
  storage,
  whitelist: ['product']
}

const persistedReducer = persistReducer(persistConfig, compareApp)

export const reduxStore = createStore(
  persistedReducer,
  {},
  compose(
      applyMiddleware(
          epicMiddleware,
          thunkMiddleware,
          routerMiddleware(browserHistory)
      )
  )
);

epicMiddleware.run(rootEpic);

export const persistor = persistStore(reduxStore);
