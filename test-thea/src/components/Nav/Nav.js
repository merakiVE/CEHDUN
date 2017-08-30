//Dependencies
import React, { Component } from 'react';

//Styles
import './Nav.css'

//Assets
import Menu from '../Menu/Menu'

//Material-UI
import Appbar from 'material-ui/AppBar';

const styles = {
	colorNav:{
	    background:'#00E676',
    }
}

class Nav extends Component{

	constructor(props) {
	  	super(props);
	  	this.state = {menu: false, modal: false};
	}

	handleToggleMenu = () => this.setState({menu: !this.state.menu});

	render(){
		return (
			<Appbar
                style={styles.colorNav}
				title={<span className={'title-nav'}>CVDI</span>}
				onLeftIconButtonTouchTap={this.handleToggleMenu}>

			    <Menu
			    	open={this.state.menu}
			    	onTouchTap={this.handleToggleMenu}
			    	onRequestChange={this.handleToggleMenu}
			    />

			</Appbar>

		);
	}
}

export default Nav;
