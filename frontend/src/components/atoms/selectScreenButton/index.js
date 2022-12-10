import React from 'react';
import styled from 'styled-components';
import { Colors } from '../../../config';

const StyledScreenButton = styled.div`
  :hover {
    background-color: ${Colors.CARD_BLUE};
  }
  display: flex;
  justify-content: center;
  align-items: center;
  height: -webkit-fill-available;
  width: -webkit-fill-available;
`;

export default StyledScreenButton;
