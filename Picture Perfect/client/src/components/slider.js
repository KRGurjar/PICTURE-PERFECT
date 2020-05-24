import React, { Component } from "react";
import Slider from "react-slick";
import styles from './App.module.css'
import 'bootstrap/dist/css/bootstrap.min.css';
import { Button , Card ,Container ,Row , Col ,Navbar, Form ,Nav ,FormControl} from 'react-bootstrap';
import Movie from  "./MovieContainer"

export default class slider extends Component {
  render() {
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
    return (
      <div>
         
        <Slider {...settings}>
          <div>
          <img src="http://image.tmdb.org/t/p/w200//wlfDxbGEsW58vGhFljKkcR5IxDj.jpg" className={styles.image}/>
          </div>
          <div>
            <img src="http://image.tmdb.org/t/p/w200//33VdppGbeNxICrFUtW2WpGHvfYc.jpg" className={styles.image}/>
          </div>
          <div>
           <img src="http://image.tmdb.org/t/p/w200//xBHvZcjRiWyobQ9kxBhO6B2dtRI.jpg" className={styles.image}/>
          </div>
          <div>
          <img src="http://image.tmdb.org/t/p/w200//c01Y4suApJ1Wic2xLmaq1QYcfoZ.jpg" className={styles.image}/>
          </div>
        </Slider>
    
      </div>
    );
  }
}



