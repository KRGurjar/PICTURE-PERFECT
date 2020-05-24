import {
	FETCH_SLIDER_REQUEST,
	FETCH_SLIDER_SUCCESS,
	FETCH_SLIDER_FAILURE
}  from "./MovieSlider.js"

const initialState={
	loading : false,
	movies  : [] ,
	error : ''
}

const reducer= (state=initialState,action) => {
	switch(action.type){
		case FETCH_SLIDER_REQUEST : return{
			...state,
			loading : true 
		}
		case FETCH_SLIDER_SUCCESS: return{
			...state,
			loading : false ,
			movies : action.payload
		}
		case FETCH_SLIDER_FAILURE : return{
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