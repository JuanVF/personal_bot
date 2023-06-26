/*
Copyright 2023 Juan Jose Vargas Fletes

This work is licensed under the Creative Commons Attribution-NonCommercial (CC BY-NC) license.
To view a copy of this license, visit https://creativecommons.org/licenses/by-nc/4.0/

Under the CC BY-NC license, you are free to:

- Share: copy and redistribute the material in any medium or format
- Adapt: remix, transform, and build upon the material

Under the following terms:

  - Attribution: You must give appropriate credit, provide a link to the license, and indicate if changes were made.
    You may do so in any reasonable manner, but not in any way that suggests the licensor endorses you or your use.

- Non-Commercial: You may not use the material for commercial purposes.

You are free to use this work for personal or non-commercial purposes.
If you would like to use this work for commercial purposes, please contact Juan Jose Vargas Fletes at juanvfletes@gmail.com.
*/
package services

import "github.com/JuanVF/personal_bot/common"

var logger *common.Logger = common.GetLogger()

var config *common.Configuration = common.GetConfig()

var googleQuery string = "from:(bncontacto@bncr.fi.cr OR popularvisa@bancopopularinforma.fi.cr OR notificacion@notificacionesbaccr.com)"

var walmartQuery string = "from:cr@mail1.gosocket.net"
