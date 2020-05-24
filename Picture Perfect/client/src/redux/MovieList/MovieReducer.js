import {
	FETCH_MOVIE_REQUEST,
	FETCH_MOVIE_SUCCESS,
	FETCH_MOVIE_FAILURE
}  from "./MovieType.js"

const initialState={
	loading : false,
	movies  : [] ,
	error : ''
}

const reducer= (state=initialState,action) => {
	switch(action.type){
		case FETCH_MOVIE_REQUEST : return{
			...state,
			loading : true 
		}
		case FETCH_MOVIE_SUCCESS: return{
			...state,
			loading : false ,
			movies : action.payload
		}
		case FETCH_MOVIE_FAILURE : return{
			...state,
			loading : false ,
			error : action.payload
		}
		default : return state
	}
}
export default reducer
export const getMovies = state => state.movies;
export const getLoading = state => state.loading;
export const getError = state => state.error;