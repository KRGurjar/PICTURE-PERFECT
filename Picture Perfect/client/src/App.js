import React, { Component } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import Home from './components/Home';
import MovieContainer from './components/MovieContainer';
import { Provider } from 'react-redux'
import store from './redux/store'
import Toolbar from './components/Toolbar';

function App() {
    return (
        <div>
        <Toolbar/>
        <Provider store={store}>
            <Switch>
                <Route path="/Movies" component={MovieContainer} exact />
                <Route path="/" component={Home}  exact />
                 <Route component={Error} />
            </Switch>
        </Provider>
        </div>
    )
}
export default App;