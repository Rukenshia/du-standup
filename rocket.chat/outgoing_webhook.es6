/* exported Script */
/* globals console, _, s, HTTP */

const service_url = '***REMOVED***';

const reMsg = /@standup (add|remove) (.*?) (.*)/i;

const requestHandlers = {
  add(data) {
    return {
      method: 'POST',
      url: `${service_url}/categories/${data.category}/entries`,
      data: {
        Title: data.title,
      },

      __meta: data,
    }
  },
  remove(data) {
    return false;
  },
};

const responseHandlers = {
  add(request, response) {
    if (response.status_code === 400) {
      return {
        content: {
          text: 'this is a little embarrassing, something went wrong. please use the [web ui](***REMOVED***) :sweat_smile:',
          parseLinks: true,
        }
      }
    } else if (response.status_code === 404) {
      return {
        content: {
          text: 'sorry, this category doesn\'t exist :sweat_smile:. try using one of `need`, `interest` or `event` :blush:',
          parseLinks: true,
        }
      }
    }

    let cat = request.__meta.category;
    if (cat.endsWith('s')) {
      cat = cat.slice(0, cat.length - 1);
    }

    return {
      content: {
        text: `alright, I added ${cat} \`${request.__meta.title}\` to the next standup :blush:`,
      }
    };
  },
  remove(request, response) {
    return false;
  }
}


/** Global Helpers
 *
 * console - A normal console instance
 * _       - An underscore instance
 * s       - An underscore string instance
 * HTTP    - The Meteor HTTP object to do sync http calls
 */

class Script {
  /**
   * @params {object} request
   */
  prepare_outgoing_request({ request }) {
    // request.params            {object}
    // request.method            {string}
    // request.url               {string}
    // request.auth              {string}
    // request.headers           {object}
    // request.data.token        {string}
    // request.data.channel_id   {string}
    // request.data.channel_name {string}
    // request.data.timestamp    {date}
    // request.data.user_id      {string}
    // request.data.user_name    {string}
    // request.data.text         {string}
    // request.data.trigger_word {string}

    const match = request.data.text.match(reMsg);

    if (!match) {
      return false;
    }

    const data = {
      event: match[1],
      category: match[2],
      title: match[3],
    };

    if (!data.category.endsWith('s')) {
      data.category = `${data.category}s`;
    }

    return requestHandlers[data.event](data);
  }

  /**
   * @params {object} request, response
   */
  process_outgoing_response({ request, response }) {
    // request              {object} - the object returned by prepare_outgoing_request

    // response.error       {object}
    // response.status_code {integer}
    // response.content     {object}
    // response.content_raw {string/object}
    // response.headers     {object}


    return responseHandlers[request.__meta.event](request, response);
  }
}