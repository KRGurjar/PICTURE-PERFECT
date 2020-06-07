import React, { Component } from "react";
import Slider from "react-slick";
import styles from './App.module.css'
import 'bootstrap/dist/css/bootstrap.min.css';
import { Button ,Form ,Nav ,FormControl} from 'react-bootstrap';
import {connect} from 'react-redux'
import {fetchMovie} from '../redux'
import { bindActionCreators } from 'redux';
import {getError,getMovies,getLoading} from '../redux/MovieList/MovieReducer'

function Square(props) {
  const value= props.value;
  const r='http://image.tmdb.org/t/p/w200/'+value.poster_path;
  const m="http://localhost:3000/Movies/" + value.id;
  return (
    <a  href={m}>
      <img src={r} className={styles.image}/>
 </a>
  );
}

export  class Home extends Component {
  constructor(props) {
        super(props);
    }
    componentDidMount() {
        const {fetchMovie} = this.props;
        fetchMovie(1,null);
      
    }

  render() {
    const { error, loading, movies } = this.props;
    const settings = {
      dots: true,
      infinite: true,
      slidesToShow: 3,
      slidesToScroll: 1,
      autoplay: true,
      speed: 2000,
      autoplaySpeed: 2000,
      cssEase: "linear"
    };
     if (error) {
      return <div>Error! {error.message}</div>;
    }

    if (loading) {
      return <div>Loading...</div>;
    }
    console.log(movies);
      return  (
          <div>
            <body>
               <div className={styles.heroimage}>
                  <div className={styles.herotext}>
                    <h1 >We Are Picture-Perfect</h1>
                     <h3>And we are here to entertain you</h3>
                     <Button>Login</Button>
                     </div>
                </div>
            </body>
                  
        <Slider {...settings}>
        {movies.map(product =>
         <div >
         <Square  value={product}  />
         </div>
        )}
          
       
        </Slider>
          <p className={styles.center} > Copyright : picture-perfect</p>
          </div>
        )
    }
}
   
const mapStateToProps = state => ({
    error: getError(state.List),
    movies: getMovies(state.List),
    loading: getLoading(state.List)
})

const mapDispatchToProps = dispatch => bindActionCreators({
    fetchMovie: fetchMovie
}, dispatch)


export default connect(
              mapStateToProps,
              mapDispatchToProps
              )(Home)   