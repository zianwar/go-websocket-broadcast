import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import './index.css';

class App extends Component {
  wsClient = null;
  state = {
    counter: 'Loading..',
    clientId: null,
    connState: null,
  };

  readyStates = {
    0: 'CONNECTING',
    1: 'OPEN',
    2: 'CLOSING',
    3: 'CLOSED'
  };

  bgs = {
    CONNECTING: 'bg-blue-500',
    OPEN: 'bg-green-500',
    CLOSING: 'bg-orange-500',
    CLOSED: 'bg-red-500'
  };

  componentDidMount = () => {
    this.wsClient = new WebSocket('ws://localhost:7777/status');

    this.wsClient.onopen = this.handleOpen;
    this.wsClient.onerror = this.handleError;
    this.wsClient.onmessage = this.handleMessage;
    this.wsClient.onclose = this.handleClose;
  };

  handleOpen = () => {
    console.log('ws connection opened');
    this.forceUpdate();
  };

  handleError = error => {
    console.error('ws error', error);
    this.forceUpdate();
  };

  handleClose = () => {
    console.error('ws connection closed');
    this.forceUpdate();
  };

  handleMessage = event => {
    const { counter, clientId } = JSON.parse(event.data);
    console.log('got message:', { counter, clientId });
    if (counter != null && clientId != null) {
      this.setState({ counter, clientId });
    }
  };

  bg = () => this.wsClient && this.bgs[this.readyStates[this.wsClient.readyState]]

  render = () => (
    <div className={`${this.bg() || ''} h-full`}>
      <div className={`flex flex-col items-center pt-48`}>
        <h1 className="text-6xl font-bold text-gray-100 uppercase tracking-wider">
          {this.state.counter}
        </h1>
        {this.state.clientId != null && (
          <div className="text-normal text-white opacity-75 uppercase text-sm tracking-wider">
            <span className="mr-1">Your ID:</span>
            <span className="font-bold">{this.state.clientId}</span>
          </div>
        )}

        <div className="w-1/2 px-5 mt-16 pt-5 pb-0 text-center text-white uppercase tracking-wider">
          <span className="mr-1 opacity-50">Connection:</span>
          <span className="font-bold">
            {(this.wsClient && this.readyStates[this.wsClient.readyState]) || 'Connecting...'}
          </span>
        </div>
      </div>
    </div>
  );
}

ReactDOM.render(<App />, document.getElementById('root'));
