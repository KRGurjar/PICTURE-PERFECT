import {
	FETCH_SLIDER_REQUEST,
	FETCH_SLIDER_SUCCESS,
	FETCH_SLIDER_FAILURE
}  from "./MovieSlider.js"
import axios from 'axios'


export const fetchSliderRequest = () =>{
	return {
	type:FETCH_SLIDER_REQUEST 
	}
}

const  fetchSliderSuccess = movies =>{
	return {
	type:FETCH_SLIDER_SUCCESS,
	payload: movies
	}
}

const fetchSliderFailure = error =>{
	return {
	type:FETCH_SLIDER_FAILURE,
	payload : error
	}
}

export const fetchSlider = () => {
	return  async (dispatch)=> {
		dispatch(fetchSliderRequest)
		var address=`http://localhost:8000/Home/slider`;
		await axios.get(address)
		.then( reponse => {
			const movies = reponse.data
			//console.log(movies)
			dispatch(fetchSliderSuccess(movies))
		})
		.catch(error => {
			const errorMsg = error.message
			dispatch(fetchSliderFailure(errorMsg))
		})

	}
}