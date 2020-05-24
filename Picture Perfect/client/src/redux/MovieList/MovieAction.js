import {
	FETCH_MOVIE_REQUEST,
	FETCH_MOVIE_SUCCESS,
	FETCH_MOVIE_FAILURE
}  from "./MovieType.js"
import axios from 'axios'


export const fetchMovieRequest = () =>{
	return {
	type:FETCH_MOVIE_REQUEST 
	}
}

const  fetchMovieSuccess = movies =>{
	return {
	type:FETCH_MOVIE_SUCCESS,
	payload: movies
	}
}

const fetchMovieFailure = error =>{
	return {
	type:FETCH_MOVIE_FAILURE,
	payload : error
	}
}

export const fetchMovie = (pageNumber,query) => {
	return  async (dispatch)=> {
		dispatch(fetchMovieRequest)
   
		var address=`http://localhost:8000/Home/${pageNumber}`;
		if(query)
          address=`http://localhost:8000/Home/${query}/${pageNumber}`;
		await axios.get(address)
		.then( reponse => {
			const movies = reponse.data
			//console.log(movies)
			dispatch(fetchMovieSuccess(movies))
		})
		.catch(error => {
			const errorMsg = error.message
			dispatch(fetchMovieFailure(errorMsg))
		})

	}
}