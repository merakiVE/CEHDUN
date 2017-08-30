//Dependencies
import React, { Component } from 'react';
import injectTapEventPlugin from 'react-tap-event-plugin';

//Styles
import './App.css';

//Assets
import Nav from './components/Nav/Nav';
import Content from './components/Content/Content';
//Material-UI
import getMuiTheme from 'material-ui/styles/getMuiTheme'
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider'

class App extends Component {

    constructor(){
        super();
       	injectTapEventPlugin();
    }

    render() {
        return (
            <MuiThemeProvider muiTheme={getMuiTheme()}>
            	<div>
            	    <Nav/>
        	        <Content body={this.props.children}/>
            	</div>
            </MuiThemeProvider> 
        );
    }
}

export default App;
