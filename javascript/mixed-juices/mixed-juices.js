// @ts-check
//
// The line above enables type checking for this file. Various IDEs interpret
// the @ts-check directive. It will give you helpful autocompletion when
// implementing this exercise.

import { time } from "console";

/**
 * Determines how long it takes to prepare a certain juice.
 *
 * @param {string} name
 * @returns {number} time in minutes
 */
export function timeToMixJuice(name) {
  switch (name) {
    case 'Pure Strawberry Joy':
      return 0.5;
    case 'Energizer':
    case 'Green Garden':
      return 1.5;
    case 'Tropical Island':
      return 3;
    case 'All or Nothing':
      return 5;
    default:
      return 2.5;


  }
}

/**
 * Calculates the number of limes that need to be cut
 * to reach a certain supply.
 *
 * @param {number} wedgesNeeded
 * @param {string[]} limes
 * @returns {number} number of limes cut
 */
export function limesToCut(wedgesNeeded, limes) {
  let limesNeeded = 0;
  let wedgesCollected = 0;
  for (let limeSize of limes) {
    let wedgeSize = 0;
    switch (limeSize) {
      case 'small':
        wedgeSize = 6;
        break;
      case 'medium':
        wedgeSize = 8;
        break;
      case 'large':
        wedgeSize = 10;
        break;
    }
    if (wedgesCollected < wedgesNeeded) {
      wedgesCollected += wedgeSize;
      limesNeeded++;
    }
  }
  return limesNeeded;
}

/**
 * Determines which juices still need to be prepared after the end of the shift.
 *
 * @param {number} timeLeft
 * @param {string[]} orders
 * @returns {string[]} remaining orders after the time is up
 */
export function remainingOrders(timeLeft, orders) {
  let ordersLeft = [];
  for(let order of orders) {
    let timeToMix = timeToMixJuice(order)
    console.log("TIME TO MIX: ", timeToMix);
    if (timeToMix <= timeLeft || timeLeft > 0) {
      timeLeft -= timeToMix;
      
    } else {
      ordersLeft.push(order);
    }
  }
  console.log("ORDERS LEFT: ", ordersLeft);
  return ordersLeft;
}
