local api = vim.api
local uv = vim.loop

local function string_to_table(str)
  -- Load the string as Lua code and execute it
  local func, err = loadstring("return " .. str)
  if err then
    print("Error parsing Lua table string:", err)
    return nil
  end

  -- Call the loaded function to obtain the Lua table
  local tbl = func()
  return tbl
end

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
    uv.read_start(child_output, function (status, data)
        vim.schedule(function ()
            if data then
                api.nvim_buf_set_lines(0, 0, -1, false, string_to_table(data) )
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
           uv.write(child_input, vim.inspect(buffer_lines).."\n")
       end)
    end)
end

local function syncerv1(room_id)
    local child_input, child_output, handle = client_setup(room_id)
    from_client(child_output)
    to_client(child_input)
end


return {syncerv1 = syncerv1}

