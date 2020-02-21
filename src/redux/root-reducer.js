import { combineReducers } from 'redux'
import { persistReducer } from 'redux-persist'
import storage from 'redux-persist/lib/storage'
import product from './product/productReducer'


const persistConfig = {
  key: 'root',
  storage,
  whitelist: ['product']
}

const compareApp = combineReducers({
  product
});

export default persistReducer(persistConfig, compareApp)
