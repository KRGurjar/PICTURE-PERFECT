
import thunk from 'redux-thunk'

import { createStore, applyMiddleware } from 'redux';
import { composeWithDevTools } from 'redux-devtools-extension';
import logger from 'redux-logger'
//import rootReducer from './rootreducer.js'
import reducer from './MovieList/MovieReducer.js'


const store = createStore (
	reducer, 
	composeWithDevTools (applyMiddleware(logger ,thunk))
	)

export default store