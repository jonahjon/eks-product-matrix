import { createStore, applyMiddleware } from 'redux';
import logger from 'redux-logger';
import { persistStore } from 'redux-persist'
import thunkMiddleware from 'redux-thunk'
import rootReducer from './root-reducer';


const middlewares = [logger, thunkMiddleware];

export const store = createStore(rootReducer, applyMiddleware(...middlewares));

export const persistor = persistStore(store)

export default { store, persistor };