import React from 'react';
import { Colors } from '../../../config';
import SelectScreenButton from '../../atoms/selectScreenButton';
import ArrowBackIcon from '@material-ui/icons/ArrowBack';
import PropTypes from 'prop-types';

const GoBackTab = (props) => {
  return (
    <div
      style={{
        color: Colors.WHITE,
        flexDirection: 'column',
        display: 'flex',
        justifyContent: 'space-evenly',
        alignItems: 'center',
        height: '6vh',
        width: '-webkit-fill-available',
      }}
    >
      <div style={{ height: '1px', backgroundColor: Colors.NAVY_BLUE, width: '-webkit-fill-available' }} />
      <SelectScreenButton style={{ height: '-webkit-fill-available' }} onClick={props.handleGoBack}>
        <ArrowBackIcon style={{ color: Colors.VALUSAFE_BLUE, marginRigth: '5px' }} />
        {/* <div style={{ fontFamily: MAIN_TEXT_FONT, color: Colors.VALUSAFE_BLUE }}>Voltar</div> */}
      </SelectScreenButton>
    </div>
  );
};

export default GoBackTab;

const { func } = PropTypes;
GoBackTab.propTypes = {
  handleGoBack: func.isRequired,
};
