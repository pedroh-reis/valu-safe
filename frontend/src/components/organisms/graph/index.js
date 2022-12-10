import React, { Component } from 'react';
import axios from 'axios';
// import { Card } from 'antd';
// import moment from 'moment';
//import { Checkbox } from 'antd';
import { LineChart, XAxis, YAxis, Tooltip, Legend, Line, CartesianGrid, ResponsiveContainer } from 'recharts';
import styled from 'styled-components';
import CircularProgress from '@mui/material/CircularProgress';
import { Grid } from '@material-ui/core';
// import Checkbox from '../components/atoms/checkbox';
import { STROKE_WIDTH, Colors, MAIN_TEXT_FONT } from '../../../config';
import PropTypes from 'prop-types';
// import unirest from 'unirest';
import apiData from '../../../mock/mock-data.json';

export default class Graph extends Component {
  state = {};

  async componentDidMount() {}

  renderColorfulLegendText = (value, entry) => {
    const { color } = entry;

    return <span style={{ color }}>{value}</span>;
  };

  handleGetFeedBack = () => {};

  render() {
    return (
      <React.Fragment>
        {/* <<<<<<<<<<<<<< LOADING >>>>>>>>>>>>> */}
        {this.props.loading && (
          <Grid container style={{ height: '72vh' }} direction="column" justify="center" alignItems="center">
            <Grid item>
              <CircularProgress color={Colors.VALUSAFE_BLUE} />
            </Grid>
          </Grid>
        )}
        {/* <<<<<<<<<<<<<< GRAPH >>>>>>>>>>>>>>> */}
        {!this.props.loading && (
          <div
            style={{
              height: '74vh',
              display: 'flex',
              flexDirection: 'column',
              alignItems: 'center',
            }}
          >
            <div
              style={{
                color: Colors.TEXT,
                marginTop: '3vh',
                marginBottom: '3vh',
                fontSize: 'medium',
                fontFamily: MAIN_TEXT_FONT,
                marginRight: '3vw',
                marginLeft: '3vw',
                textAlign: 'center',
                fontWeight: 'bold',
              }}
            >
              {`Atividade do Locker ${this.props.lockerId + 1} - 24h`}
            </div>
           

            <ResponsiveContainer width="90%" height="70%">
              <LineChart
                width="80vw"
                height="70vh"
                data={this.props.data}
                margin={{ top: 10, right: 30, left: 20, bottom: 5 }}
                style={{
                  backgroundColor: Colors.BACKGROUND,
                }}
              >
                <CartesianGrid strokeDasharray="3 3" />
                <XAxis dataKey="time" style={{ fontSize: 'small' }} />
                <YAxis type="category" style={{ fontSize: 'small', marginLeft: '-' }} />
                <Tooltip
                  wrapperStyle={{ borderColor: 'blue' }}
                  contentStyle={{
                    backgroundColor: Colors.BACKGROUND,
                    borderRadius: 10,
                  }}
                  labelStyle={{ color: Colors.TEXT }}
                />
                <Legend formatter={this.renderColorfulLegendText} />
                <Line
                  type="stepAfter"
                  dataKey="status"
                  dot={true}
                  stroke={Colors.VALUSAFE_BLUE}
                  strokeWidth={STROKE_WIDTH}
                />
              </LineChart>
            </ResponsiveContainer>
          </div>
        )}
      </React.Fragment>
    );
  }
}
