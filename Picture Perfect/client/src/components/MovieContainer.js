import React , {Component} from 'react'
import {connect} from 'react-redux'
import {fetchMovie} from '../redux'
import { bindActionCreators } from 'redux';
import {getError,getMovies,getLoading} from '../redux/MovieList/MovieReducer'


import Pagination from "react-js-pagination";
import 'bootstrap/dist/css/bootstrap.min.css';
import { Button , Card ,Container ,Row , Col ,Navbar, Form ,Nav ,FormControl} from 'react-bootstrap';
import styles from './App.module.css'


function Square(props) {
  const value= props.value;
  const r='http://image.tmdb.org/t/p/w154/'+value.poster_path;
  return (
    <a  href="http://localhost:3000/Movies">
    <Card border="dark" style={{ width: '26rem'}} >
  <Card.Img variant="top" src={r}  />
  <Card.Body>
   <h4 style={{backgroundColor: "lightblue"}} >{value.title}</h4>
  </Card.Body>
</Card> 
 </a>
  );
}

class MovieContainer extends Component {
    constructor(props) {
        super(props);
        this.handleSubmit = this.handleSubmit.bind(this)
 
        this. state = { users: null,
    total: null,
    per_page: null,
    activePage: 1 ,
   query:''}
    }

    componentDidMount() {
        const {fetchMovie} = this.props;
        fetchMovie(1,this.state.query);
        
      
    }
   
   handleInputChange = event => {
    const query = event.target.value;
    this.setState({query:query});
    const {fetchMovie} = this.props;
    fetchMovie(1,query);
    //console.log(query);
  }

     handlePageChange(pageNumber) {
    console.log(`active page is ${pageNumber}`);
    this.setState({activePage: pageNumber});
   const {fetchMovie} = this.props;
        fetchMovie(pageNumber,this.state.query);
  }
     handleSubmit() {
    //console.log("sub",this.state.query);
   const {fetchMovie} = this.props;
        fetchMovie(1,this.state.query);
  }

    render() {
    const { error, loading, movies } = this.props;
    
    if (error) {
      return <div>Error! {error.message}</div>;
    }

    if (loading) {
      return <div>Loading...</div>;
    }


    return (
    	<div  className={styles.all} >
      <main style={{marginTop:'3px'}}  className={styles.center}>
       <Form className={styles.form} >
    <Form.Control 
          type="text" 
          value={this.state.query}
          onChange={this.handleInputChange}
          placeholder="Search Movie" 
          required  />
   <Button variant="primary"  onClick={this.handleSubmit} >
    Submit
  </Button  >
</Form>
</main>
      <div>
      {movies ? (
        <div>
        <div className={styles.center}>
               <div className={styles.gridcontainer}>
        {movies.map(product =>
         <div className={styles.item}>
         <Square  value={product}  />
         </div>
        )}
         </div>
         </div>
                <div  className={styles.center}>
                 <Pagination
                  size="lg"
          className={styles.page}
          activePage={this.state.activePage}
          itemsCountPerPage={10}
          totalItemsCount={90}
          pageRangeDisplayed={5}
          onChange={this.handlePageChange.bind(this)}
        />
      </div>
      </div>
      ) : (
         <h3  className={styles.center}> no results </h3> 
      )}
    </div>
        
      </div>

    );
    
  }
}


const mapStateToProps = state => ({
    error: getError(state),
    movies: getMovies(state),
    loading: getLoading(state)
})

const mapDispatchToProps = dispatch => bindActionCreators({
    fetchMovie: fetchMovie
}, dispatch)


export default connect(
              mapStateToProps,
              mapDispatchToProps
              )(MovieContainer)