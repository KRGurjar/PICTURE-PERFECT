import { combineReducers } from 'redux'
import reducer from './MovieList/MovieReducer.js'

const rootReducer = combineReducers ({
	 movies : reducer
	)}
 
 export defaul rootReducer