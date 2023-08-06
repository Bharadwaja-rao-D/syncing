local api = vim.api
local uv = vim.loop

local utils = require('nvim-syncer.utils')

local function client_setup(room_id)
    local cmd = "syncing"
    local child_input = uv.new_pipe()
    local child_output = uv.new_pipe()
    local child_err = uv.new_pipe()
    local options = {args = {room_id}, stdio = {child_input, child_output, child_err}}
    local handle;

    -- TODO: proper closing should be done
    local on_exit = function (status)
        print(status)
        uv.close(handle)
        uv.close(child_input)
        uv.close(child_err)
    end

    handle = uv.spawn(cmd, options, on_exit)

    return child_input, child_output, handle
end

local function from_client(child_output)
    local buffer = api.nvim_get_current_buf()
    uv.read_start(child_output, function (status, data)
        vim.schedule(function ()
            if data then
                data = '{"start_index":0,"end_index": -1, "text":["line1", "line2"]}'
                data = utils.deser(data)
                print(data.start_index)
                api.nvim_buf_set_lines(buffer, tonumber(data.start_line), tonumber(data.end_line), false, data.text)
            end
        end)
    end)
end

local function to_client(child_input)
    local buffer = api.nvim_get_current_buf()
    local timer = uv.new_timer()
    timer:start(1000, 7000, function ()
       vim.schedule(function ()
           local buffer_lines = api.nvim_buf_get_lines(buffer,0,-1,false)
           uv.write(child_input, utils.ser({start_index = 0, end_index = -1, text = buffer_lines}))
       end)
    end)
end

local function syncerv1(room_id)
    local child_input, child_output = client_setup(room_id)
    from_client(child_output)
    to_client(child_input)
end


return {syncerv1 = syncerv1}
